package test;

import java.util.ArrayList; 
import java.util.Scanner;

public class TodoList {
    public static void main(String[] args) {
        ArrayList<String> tasks = new ArrayList<>();
        Scanner scanner = new Scanner(System.in);

        while (true) {
            System.out.println("\n===== Todolist =====");
            System.out.println("1. Add task");
            System.out.println("2.Display task list");
            System.out.println("3.completed task");
            System.out.println("4.Exit");
            System.out.println("please select (1~4)");
            
            int choice = scanner.nextInt();
            scanner.nextLine();

            switch (choice) {
                case 1:
                    System.out.print("please enter a new task: ");
                    String newTask = scanner.nextLine();
                    tasks.add(newTask);
                    System.out.println("Task added!");
                    break;

                case 2:
                    if (tasks.isEmpty()) {
                        System.out.println("\nThere are no current tasks... ");
                    } else{
                        System.out.println("\nCurrent tasks: ");
                        for (int i = 0; i < tasks.size(); i++) {
                            System.out.println((i + 1) + ":" + tasks.get(i));
                        }
                    }
                    break;
                
                case 3:
                    if (tasks.isEmpty()) {
                        System.out.println("There are no tasks");
                    } else {
                        System.out.println("please enter the number of the task to completed : ");
                        int taskNum = scanner.nextInt();
                        if (taskNum > 0 && taskNum <= tasks.size()) {
                            String completedTask = tasks.remove(taskNum - 1);
                            System.out.println("task「" + completedTask + "」 completed!");
                        } else {
                            System.out.println("Invalid task number");
                        }
                    }
                    break;

                case 4:
                    System.out.println("the application is closed");
                    scanner.close();
                    return;

                default:
                    System.out.println("Invalid selection. Please enter a number between 1 and 4.");
            }
        }
    }
}