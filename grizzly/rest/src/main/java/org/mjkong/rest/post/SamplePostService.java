package org.mjkong.rest.post;

import org.jvnet.hk2.annotations.Service;
import org.mjkong.rest.api.TestObject;

@Service
public class SamplePostService implements PostService {
    @Override
    public TestObject test(String input) {
        return TestObject.builder().testProp(input).build();
    }
}