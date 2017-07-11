--
-- Created by IntelliJ IDEA.
-- User: admin
-- Date: 2017/7/11
-- Time: 15:18
-- To change this template use File | Settings | File Templates.
--

function heroLevel(a)
    print("英雄id:" .. a:heroId() .. ",level:" .. a:heroLv() .. ",star:" .. a:heroStar())
    return a:heroLv() * 10
end