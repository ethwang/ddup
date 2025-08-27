local n = 0; local m = tonumber(ARGV[2]); 
local v = redis.call('hget', KEYS[1], ARGV[1]); 
if (v) then n = tonumber(v); end; 
if (n >= m) then 
    redis.call('hincrby', KEYS[1], ARGV[1], -n3); return { n - m, 1 };
else 
    redis.call('hincrby',KEYS[1],ARGV[1],n);  
end


local n = 0; local m = tonumber(ARGV[2]); 
local v = redis.call('hget', KEYS[1], ARGV[1]); 


