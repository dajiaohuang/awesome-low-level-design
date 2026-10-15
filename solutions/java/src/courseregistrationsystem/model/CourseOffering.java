package courseregistrationsystem.model;

import courseregistrationsystem.observer.CourseOfferingObserver;

import java.util.*;

public class CourseOffering {
    private final String id;
    private final Course course;
    private final Professor professor;
    private final TimeSlot timeSlot;
    private final int capacity;
    private final List<Student> registeredStudents = new ArrayList<>();
    private final Queue<Student> waitlistedStudents = new LinkedList<>();
    private final List<CourseOfferingObserver> observers = new ArrayList<>();

    private CourseOffering(Builder builder) {
        this.id = builder.id;
        this.course = builder.course;
        this.professor = builder.professor;
        this.timeSlot = builder.timeSlot;
        this.capacity = builder.capacity;
    }

    public void addObserver(CourseOfferingObserver observer) { observers.add(observer); }
    private void notifyObservers() { observers.forEach(o -> o.onSpotAvailable(this)); }

    public synchronized boolean isFull() { return registeredStudents.size() >= capacity; }
    public synchronized void addStudent(Student student) { registeredStudents.add(student); student.enroll(this); }
    public synchronized void addToWaitlist(Student student) { waitlistedStudents.add(student); }

    public synchronized boolean registerOrWaitlist(Student student) {
        if (isFull()) {
            waitlistedStudents.add(student);
            return false;
        }
        registeredStudents.add(student);
        student.enroll(this);
        return true;
    }

    public void dropStudent(Student student) {
        boolean spotAvailable;
        synchronized (this) {
            boolean wasFull = isFull();
            registeredStudents.remove(student);
            student.drop(this);
            spotAvailable = wasFull && !isFull();
        }
        if (spotAvailable) {
            notifyObservers();
        }
    }

    public synchronized Optional<Student> getNextFromWaitlist() { return Optional.ofNullable(waitlistedStudents.poll()); }
    public String getId() { return id; }
    public Course getCourse() { return course; }
    public TimeSlot getTimeSlot() { return timeSlot; }

    // Builder Pattern
    public static class Builder {
        private String id; private Course course; private Professor professor; private TimeSlot timeSlot; private int capacity;
        public Builder(String id, Course course) { this.id = id; this.course = course; }
        public Builder withProfessor(Professor professor) { this.professor = professor; return this; }
        public Builder at(TimeSlot timeSlot) { this.timeSlot = timeSlot; return this; }
        public Builder withCapacity(int capacity) { this.capacity = capacity; return this; }
        public CourseOffering build() { return new CourseOffering(this); }
    }
}
