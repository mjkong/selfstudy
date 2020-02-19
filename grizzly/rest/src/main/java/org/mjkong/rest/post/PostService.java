package org.mjkong.rest.post;

import org.jvnet.hk2.annotations.Contract;
import org.mjkong.rest.api.TestObject;

@Contract
public interface PostService {
    TestObject test(String input);
}