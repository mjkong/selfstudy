package org.mjkong.rest.post;

import javax.inject.Inject;

import org.jvnet.hk2.annotations.Service;
import org.mjkong.rest.api.TestObject;
import org.mjkong.rest.api.TestResource;

@Service
public class SampleTestResource implements TestResource {

    private PostService postService;

    @Inject
    public SampleTestResource(PostService postService) {
        this.postService = postService;
    }

    public String checkHealth(String value) {
        return postService.test(value).getTestProp();
    }


    public TestObject post(TestObject testObject) {
        return postService.test(testObject.getTestProp());
    }

    @Override
    public String checkHealth() {
      return null;
    }
}