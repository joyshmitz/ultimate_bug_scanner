import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.io.FileInputStream;
import java.io.IOException;
import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;

public class ResourceLifecycle {
    public void leak() {
        ExecutorService exec = Executors.newSingleThreadExecutor();
        exec.submit(() -> System.out.println("work"));
        // missing exec.shutdown()
    }

    public void leakStream() throws IOException {
        FileInputStream in = new FileInputStream("/tmp/data.txt");
        System.out.println(in.read());
        // missing in.close()
    }

    public void leakJdbc(Connection conn) throws SQLException {
        Statement stmt = conn.createStatement();
        PreparedStatement ps = conn.prepareStatement("SELECT * FROM users WHERE id = ?");
        ps.setInt(1, 42);
        ResultSet rs = stmt.executeQuery("SELECT NOW()");
        System.out.println(rs);
        // missing stmt.close(), ps.close(), rs.close()
    }
}
