package com.mszlu.blog.service;

import com.mszlu.blog.vo.Result;
import com.mszlu.blog.vo.params.LoginParams;

public interface LoginService {
    Result login(LoginParams loginParams);
}
