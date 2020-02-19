package org.mjkong.rest;

import java.io.IOException;
import java.net.URI;

import org.glassfish.grizzly.http.server.HttpServer;
import org.glassfish.hk2.api.ServiceLocator;
import org.glassfish.hk2.utilities.ServiceLocatorUtilities;
import org.glassfish.jersey.grizzly2.httpserver.GrizzlyHttpServerFactory;

import lombok.extern.slf4j.Slf4j;

@Slf4j
public class Main {

  public static void main(String[] args) throws Exception {
    log.info("Test");

    String BASE_URI = "http://0.0.0.0:9005/microblog/";
    ServiceLocator locator = ServiceLocatorUtilities.createAndPopulateServiceLocator();

    HttpServer httpServer = GrizzlyHttpServerFactory.createHttpServer(URI.create(BASE_URI), new JerseyApplication(), locator);

    try{

      httpServer.start();
      System.out.println(String.format("Jersey app started.\nHit enter to stop it...", BASE_URI));
      System.in.read();
    }catch(IOException e){
      log.error("Error starting server: " + e.getLocalizedMessage());
    }
  }
}