package org.mjkong.rest.api;

import javax.inject.Singleton;
import javax.ws.rs.*;

import org.jvnet.hk2.annotations.Contract;

@Contract
@Path("test")
@Singleton
public interface TestResource {

  @GET
  @Produces("text/plain")
  @Path("{value}")
  String getValue(@PathParam("value") String value);

  @POST
  @Produces("application/json")
  TestObject post(TestObject testObject) ;

}