package atm;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.CountDownLatch;

public final class ATMSystemSingletonTest {
    public static void main(String[] args) throws InterruptedException {
        int callers = 64;
        CountDownLatch ready = new CountDownLatch(callers);
        CountDownLatch start = new CountDownLatch(1);
        List<ATMSystem> instances = new ArrayList<>();
        List<Thread> threads = new ArrayList<>();
        for (int i = 0; i < callers; i++) {
            Thread thread = new Thread(() -> {
                ready.countDown();
                try {
                    start.await();
                } catch (InterruptedException e) {
                    Thread.currentThread().interrupt();
                    throw new AssertionError(e);
                }
                synchronized (instances) {
                    instances.add(ATMSystem.getInstance());
                }
            });
            threads.add(thread);
            thread.start();
        }
        ready.await();
        start.countDown();
        for (Thread thread : threads) {
            thread.join();
        }
        ATMSystem expected = instances.get(0);
        if (instances.stream().anyMatch(instance -> instance != expected)) {
            throw new AssertionError("concurrent callers received distinct ATM systems");
        }
    }
}
