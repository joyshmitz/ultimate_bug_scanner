import java.io.ByteArrayInputStream;
import java.io.InputStreamReader;
import java.lang.reflect.Field;
import java.lang.reflect.Method;
import java.math.BigDecimal;
import java.nio.file.Paths;
import java.sql.CallableStatement;
import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.Statement;
import java.util.Hashtable;
import java.util.Optional;
import java.util.Random;
import java.util.Vector;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

public final class AstGrepRulePackCoverage {
    private final Logger logger = new Logger();

    public void exercise(
            String name,
            String base,
            String child,
            String input,
            ByteArrayInputStream bytes,
            Connection existingConnection,
            Object receiver)
            throws Exception {
        Exception error = new Exception("boom");
        error.printStackTrace();
        System.out.println("debug " + name);

        Optional<String> maybe = Optional.of(name);
        if (maybe.isPresent()) {
            String present = maybe.get();
            logger.info("present " + present);
        }
        if (!maybe.isEmpty()) {
            logger.debug("not empty");
        }
        String fallback = maybe.orElse(null);
        logger.warn("fallback " + fallback);

        Paths.get(base + "/" + child);
        Paths.get(base, child + ".tmp");
        String apiToken = "hardcoded-token";
        if ("root" == name) {
            logger.info("token " + apiToken);
        }

        BigDecimal amount = new BigDecimal("1.0");
        amount.equals(new BigDecimal("1.00"));

        synchronized (this) {
            java.lang.Thread.sleep(1);
        }
        Object monitor = new Object();
        synchronized (monitor) {
            monitor.notify();
        }

        new Thread(() -> logger.info("manual")).start();
        ExecutorService cached = Executors.newCachedThreadPool();
        java.util.concurrent.ExecutorService executor =
                java.util.concurrent.Executors.newSingleThreadExecutor();
        Thread worker = new Thread(() -> logger.info("worker"));
        worker.start();

        java.sql.Connection leaked =
                java.sql.DriverManager.getConnection("jdbc:sqlite::memory:");
        ResultSet rs = existingConnection.createStatement().executeQuery("SELECT 1");
        Statement stmt = existingConnection.createStatement();
        PreparedStatement prepared = existingConnection.prepareStatement("SELECT ?");
        CallableStatement callable = existingConnection.prepareCall("{ call refresh_user(?) }");
        logger.info(String.valueOf(rs));
        logger.info(String.valueOf(stmt));
        logger.info(String.valueOf(prepared));
        logger.info(String.valueOf(callable));

        new Random();
        javax.net.ssl.HttpsURLConnection.setDefaultHostnameVerifier((host, session) -> true);
        java.security.MessageDigest.getInstance("MD5");
        String url = "http://example.test/path";
        logger.info(url);
        new java.io.ObjectInputStream(bytes).readObject();
        new InputStreamReader(bytes);
        new String(bytes.readAllBytes());
        name.getBytes();
        Vector<String> vector = new Vector<>();
        Hashtable<String, String> table = new Hashtable<>();
        logger.info(vector.toString() + table.toString());

        java.util.List.of(name).parallel().forEach(item -> logger.info("item " + item));
        Class<?> clazz = Class.forName(name);
        Field field = clazz.getDeclaredField("value");
        Method method = clazz.getDeclaredMethod("toString");
        method.invoke(receiver);
        field.setAccessible(true);
        String dangerousRegex = "(a+)+";
        logger.info(dangerousRegex);
        java.lang.Thread.ofVirtual().start(() -> logger.info("virtual"));
        java.io.FileInputStream stream = new java.io.FileInputStream(input);
        logger.info(String.valueOf(stream));
        logger.error("name " + name);
        logger.debug("cached " + cached);
        logger.debug("leaked " + leaked);
    }

    private static final class Logger {
        void debug(String message) {}
        void info(String message) {}
        void warn(String message) {}
        void error(String message) {}
    }
}
