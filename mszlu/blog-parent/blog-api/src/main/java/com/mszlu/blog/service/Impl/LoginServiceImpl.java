package com.mszlu.blog.service.Impl;

import com.mszlu.blog.service.LoginService;
import com.mszlu.blog.vo.Result;
import com.mszlu.blog.vo.params.LoginParams;
import org.apache.commons.lang3.StringUtils;
import org.springframework.stereotype.Service;

@Service
public class LoginServiceImpl implements LoginService {
    @Override
    public Result login(LoginParams loginParams) {
        /**
         * 检查登录是否合法
         * 根据用户名和密码去user表中去查询 是否存在
         * 如果不粗不在 登录失败
         * 如果存在 使用JWT 生成token ： user信息，设置过期时间
         * （登录认证的时候 先认证token字符串是否合法，去redis认证是否存在）
         */
        String account = loginParams.getAccount();
        String password = loginParams.getPassword();
        if (StringUtils.isBlank(account) || StringUtils.isBlank(password)){

        }
        return null;
    }
}
