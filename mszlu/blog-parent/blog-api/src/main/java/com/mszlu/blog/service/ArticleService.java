package com.mszlu.blog.service;

import com.mszlu.blog.vo.Result;
import com.mszlu.blog.vo.params.PageParams;
import org.springframework.stereotype.Service;

@Service
public interface ArticleService {

    Result listArtricle(PageParams pageParams);

    Result newArticles(int limit);

    Result listArchives();

    Result hotArticle(int limit);
}
