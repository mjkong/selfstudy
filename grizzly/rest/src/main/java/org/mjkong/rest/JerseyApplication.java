package org.mjkong.rest;

import com.fasterxml.jackson.jaxrs.json.JacksonJaxbJsonProvider;

import org.glassfish.hk2.utilities.binding.AbstractBinder;
import org.glassfish.jersey.server.ResourceConfig;
import lombok.extern.slf4j.Slf4j;

@Slf4j
public class JerseyApplication extends ResourceConfig {

  public JerseyApplication() {
    log.info("setting up hk2");
    packages("org.mjkong.rest", "org.mjkong.rest.api", "org.mjkong.rest.jersey");

    JacksonJaxbJsonProvider jacksonJaxbJsonProvider = new JacksonJaxbJsonProvider();
    jacksonJaxbJsonProvider.setMapper(new ObjectMapperFactory().buildObjectMapper());
    register(jacksonJaxbJsonProvider);
    register(new AbstractBinder() {
        @Override
        protected void configure() {

            // Shop to manually bind objects, in the case that the Jersey Auto-scan isn't working
            // e.g. bind(x.class).to(y.class).in(Singleton.class);
            // e.g. bind(x.class).to(y.class);
            //
            // note: if the object is generic, use TypeLiteral
            // e.g. bind(x.class).to(new TypeLiteral&lt;InjectionResolver&gt;(){});
            //
        }
    });
  }
}