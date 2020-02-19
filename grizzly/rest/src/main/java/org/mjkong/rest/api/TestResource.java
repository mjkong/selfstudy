package org.mjkong.rest.api;

import javax.inject.Singleton;
import javax.ws.rs.GET;
import javax.ws.rs.POST;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;

import org.jvnet.hk2.annotations.Contract;

@Contract
@Path("test")
@Singleton
public interface TestResource {

  @GET
  @Produces("text/plain")
  String checkHealth();

  @POST
  @Produces("application/json")
  TestObject post(TestObject testObject) ;

}