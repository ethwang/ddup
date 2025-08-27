local ngx = ngx
local core = require("apisix.core")
local protoc = require("protoc")

local schema = {type = "object"}

local plugin_name = "micolite-auth"

local _M = {
    version = 0.1,
    priority = 11003,
    name = plugin_name,
    schema = schema
}

function _M.check_schema(conf, schema_type)
    return core.schema.check(schema, conf)
end

function _M.init() end

function _M.destroy() end

function _M.rewrite(conf, ctx)
    local red_c = connect_redis()

    local key_set = {
        "did", "afid", "sys-version", "vc", "vn", "locale", "tz-secs", "mcc",
        "pkg", "os", "lang", "model", "fbid", "uid", "show-id", "token"
    }

    local evalstr =
        "local x = redis.call('zscore',KEYS[1],ARGV[1]); if (not x) then return nil end; local r = redis.call('get', KEYS[2]); if (not r) then return nil; end redis.call('zadd',KEYS[1],ARGV[2],ARGV[1]) return r"

    local raw_header = ngx.req.get_headers()
    for k, v in pairs(raw_header) do core.log.info("sugo k:", k, " v:", v) end

    local http_version = ngx.req.http_version()

    local s_cookie = ngx.var.http_cookie
    local raw_cookie
    if s_cookie then raw_cookie = get_cookie(s_cookie) end

    local uid
    local did
    local token

    for i = 1, #key_set do
        local key = key_set[i]
        local value = raw_header[key]

        if not value and raw_cookie then
            value = raw_cookie[key]
            ngx.req.set_header(key, ngx.escape_uri(value))
        end

        if key == "uid" then
            uid = value
        elseif key == "did" then
            did = value
        elseif key == "token" then
            token = value
        end
    end

    if not token then token = raw_header["key"] end

    local key3 = "device_ban"

    if did then
        local status, err1 = red_c:sismember(key3, did)
        if err1 then
            core.log.err("red_c device_ban failed:", err1)
        else
            if status and status ~= ngx.null then
                ngx.req.set_header("device_status", ngx.escape_uri(status))
            else
                core.log.err("redis_cluster device_ban, device_status=", status)
            end
        end
    end
    if ((not uid) and (not token)) then
        ngx.req.clear_header("sgu")
    else
        if ((not token)) then
            ngx.log(ngx.ERR, "micolite uid=", uid, " without token")
            core.response.add_header("Content-Type", "application/grpc")
            core.response.add_header("grpc-status", "16")
            core.response.add_header("grpc-message", "UNAUTHENTICATED2")
            return core.response.exit(401)
        elseif (uid) then
            local key1 = "accountapi_uid_2_token:{" .. uid .. "}"
            local key2 =
                "accountapi_token_2_uid:{" .. uid .. "}:" .. did .. ":" .. token
            local res, err = red_c:eval(evalstr, 2, key1, key2,
                                        did .. ":" .. token,
                                        tostring(ngx.time() * 1000))
            if err then
                ngx.log(ngx.ERR, "red_c failed:", err, "key1=", key1, "key2=",
                        key2)
                core.response.add_header("Content-Type", "application/grpc")
                core.response.add_header("grpc-status", "13")
                core.response.add_header("grpc-message", "INTERNAL")
                return core.response.exit(401)
            else
                ngx.req.clear_header("sgu")
                if res and res ~= ngx.null then
                    ngx.req.set_header("sgu", ngx.escape_uri(res))
                else
                    if http_version == 2 then
                        ngx.log(ngx.ERR, "redis_cluster empty key1=", key1,
                                ",key2=", key2, ", res=", res)
                        core.response.add_header("Content-Type",
                                                 "application/grpc")
                        core.response.add_header("grpc-status", "16")
                        core.response.add_header("grpc-message",
                                                 "UNAUTHENTICATED")
                        return core.response.exit(401)
                    else
                        return 403
                    end
                end
            end
        end
    end
end

function get_cookie(s_cookie)
    local cookie = {}
    for item in string.gmatch(s_cookie, "[^;]+") do
        local _, _, k, v = string.find(item, "^%s*(%S+)%s*=%s*(%S+)%s*")
        if k ~= nil and v ~= nil then cookie[k] = v end
    end
    return cookie
end

function connect_redis()
    local config = {
        name = "ProdCluster",
        serv_list = {
            {
                ip = "r-t4n0w4mbkt95pk5iyg.redis.singapore.rds.aliyuncs.com",
                port = 6379
            }
        },
        keepalive_timeout = 60000,
        keepalive_cons = 200,
        connect_timeout = 1000,
        read_timeout = 1000,
        send_timeout = 1000,
        max_redirection = 5,
        max_connection_attempts = 1,
        auth = "SFyVkJXsbJGV3K9D"
    }
    local redis_cluster = require "resty.rediscluster"
    return redis_cluster:new(config)
end

function _M.log(conf, ctx) end

function _M.access(conf, ctx) end

return _M
