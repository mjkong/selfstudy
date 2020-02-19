package org.cobbsalad.grizzly.client;

import org.glassfish.grizzly.Connection;
import org.glassfish.grizzly.filterchain.FilterChainBuilder;
import org.glassfish.grizzly.filterchain.TransportFilter;
import org.glassfish.grizzly.nio.transport.TCPNIOTransport;
import org.glassfish.grizzly.nio.transport.TCPNIOTransportBuilder;
import org.glassfish.grizzly.utils.StringFilter;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.nio.charset.Charset;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.Future;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.TimeoutException;
import java.util.logging.Logger;

public class EchoClient {

    private static final Logger logger = Logger.getLogger(EchoClient.class.getName());

    public static void main(String[] args) throws IOException, InterruptedException, ExecutionException, TimeoutException {


        Connection connection = null;

        FilterChainBuilder filterChainBuilder = FilterChainBuilder.stateless();
        filterChainBuilder.add(new TransportFilter());

        filterChainBuilder.add(new StringFilter(Charset.forName("UTF-8")));

        filterChainBuilder.add(new ClientFilter());

        final TCPNIOTransport transport = TCPNIOTransportBuilder.newInstance().build();
        transport.setProcessor(filterChainBuilder.build());

        try{
            transport.start();

            Future<Connection> future = transport.connect("localhost", 7777);
            connection = future.get(10, TimeUnit.SECONDS);

            assert connection != null;

            System.out.println("Ready... (\"q\" to exit)");
            final BufferedReader inReader = new BufferedReader(new InputStreamReader(System.in));
            do{
                final String userInput = inReader.readLine();
                if (userInput == null || "q".equals(userInput)){
                    break;
                }

                connection.write(userInput);
            } while (true);
        } finally {
            if(connection != null) {
                connection.close();
            }

            transport.shutdownNow();
        }
    }
}
