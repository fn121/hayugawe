import { Suspense } from "react";

interface Task {
  id: string;
  title: string;
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
  return <li>{task.title}</li>;
}

export default function TasksPage() {
  return (
    <Suspense fallback={<p>Memuat data....</p>}>
      <TaskList />
    </Suspense>
  );
}
