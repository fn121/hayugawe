import { refresh } from "next/cache";
import { Suspense } from "react";
import { CheckButton } from "./check-button";

export interface Task {
  id: string;
  title: string;
  is_done: boolean;
}

export interface PostTask {
  title: string;
}

export async function createTask(formData: FormData) {
  "use server";
  const title = formData.get("title");
  const postTask: PostTask = {
    title: title ? title.toString() : "untitled task",
  };

  // @TODO: handle not ok status response
  await fetch("http://localhost:8080/api/tasks", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(postTask),
  });

  refresh();
}

async function TaskList() {
  const data = await fetch("http://localhost:8080/api/tasks");
  const tasks: Task[] = await data.json();
  return (
    <ul>
      {tasks.map((task) => (
        <TaskItem key={task.id} task={task} />
      ))}
    </ul>
  );
}

function TaskItem({ task }: { task: Task }) {
  return (
    <li>
      <CheckButton task={task} />
      <span>{task.title}</span>
    </li>
  );
}

function FormAddTask() {
  return (
    <form action={createTask}>
      <input type="text" name="title" />
      <button type="submit">Add</button>
    </form>
  );
}

export default function TasksPage() {
  return (
    <div>
      <FormAddTask />
      <Suspense fallback={<p>Memuat data....</p>}>
        <TaskList />
      </Suspense>
    </div>
  );
}
