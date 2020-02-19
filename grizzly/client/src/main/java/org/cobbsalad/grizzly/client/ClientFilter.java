package org.cobbsalad.grizzly.client;

import org.glassfish.grizzly.filterchain.BaseFilter;
import org.glassfish.grizzly.filterchain.FilterChainContext;
import org.glassfish.grizzly.filterchain.NextAction;

import java.io.IOException;

public class ClientFilter extends BaseFilter {

    @Override
    public NextAction handleRead(FilterChainContext ctx) throws IOException {
        final String serverResponse = ctx.getMessage();

        System.out.println("Server echo: " + serverResponse);

        return ctx.getStopAction();
    }
}
