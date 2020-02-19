package org.cobbsalad.grizzly.server;

import org.glassfish.grizzly.filterchain.FilterChainBuilder;
import org.glassfish.grizzly.filterchain.TransportFilter;
import org.glassfish.grizzly.nio.transport.TCPNIOTransport;
import org.glassfish.grizzly.nio.transport.TCPNIOTransportBuilder;
import org.glassfish.grizzly.utils.StringFilter;

import java.io.IOException;
import java.nio.charset.Charset;
import java.util.logging.Logger;

public class EchoServer {

    private static final Logger logger = Logger.getLogger(EchoServer.class.getName());

    public static final String HOST = "localhost";
    public static final int PORT = 7777;


    public static void main(String[] args) throws IOException {

        FilterChainBuilder filterChainBuilder = FilterChainBuilder.stateless();
        filterChainBuilder.add(new TransportFilter());

        filterChainBuilder.add(new StringFilter(Charset.forName("UTF-8")));

        filterChainBuilder.add(new EchoFilter());

        final TCPNIOTransport transport = TCPNIOTransportBuilder.newInstance().build();

        transport.setProcessor(filterChainBuilder.build());
        try{
            transport.bind(HOST, PORT);

            transport.start();

            logger.info("Press any key to stop the server...");
            System.in.read();
        }finally{
            logger.info("Stopping transport...");
            transport.shutdownNow();


            logger.info("Stopped transport...");
        }
    }
}
