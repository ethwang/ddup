local m = tonumber(ARGV[1]);
local total = 0;
for i,arg in pairs(ARGV) do
    if (i~=1) then
        local n = 0;
        local v = redis.call('hget', KEYS[1], arg);
        if (v) then
            n = tonumber(v);
        end
        total = total+n
    end
end
if (total < m) then
    return {0,0};
end
for i,arg in pairs(ARGV) do
    if (i~=1) then
        local n = 0;
        local v = redis.call('hget', KEYS[1], arg);
        if (v) then
            n = tonumber(v);
        end
        if (n >= m) then
            redis.call('hincrby', KEYS[1], arg, -m); return {n-m,1};
        else
            redis.call('hincrby', KEYS[1], arg, -n);
            m = m-n;
        end
    end
end
return {0,0};