package com.mszlu.blog.service;

import com.mszlu.blog.dao.pojo.SysUser;
import com.mszlu.blog.vo.UserVo;

public interface SysUserService {
    SysUser findUserById(Long id);
}
