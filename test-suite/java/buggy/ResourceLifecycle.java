import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

public class ResourceLifecycle {
    public void leak() {
        ExecutorService exec = Executors.newSingleThreadExecutor();
        exec.submit(() -> System.out.println("work"));
        // missing exec.shutdown()
    }
}
