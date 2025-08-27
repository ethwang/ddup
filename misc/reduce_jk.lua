local total_coins = redis.call('get', KEYS[1]);
local reduce_amount = ARGV[1];
local times = ARGV[3];
if tonumber(total_coins) < tonumber(ARGV[4]) then return {-1, total_coins}; end;
if tonumber(total_coins) < tonumber(reduce_amount) then return -1; end;
redis.call('INCRBYFLOAT',KEYS[1],-reduce_amount);
local after_coins = redis.call('get', KEYS[1]); return {times, after_coins};


--[[

    val1 = 奖池减少的金币
    val2 = 
    val3 = 倍率 
    val4 = 奖池开奖金额
	keys := []string{jackpotKey}
	args := []string{strconv.FormatInt(time*price, 10), strconv.FormatInt(price, 10), strconv.FormatInt(time, 10), strconv.FormatInt(model.G_cfg.BonusCoins, 10)}
]]