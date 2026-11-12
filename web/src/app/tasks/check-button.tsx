"use client";

import { markTaskAsDone, markTaskAsTodo } from "./actions";
import { Task } from "./page";

export function CheckButton({ task }: { task: Task }) {
  return (
    <input
      type="checkbox"
      checked={task.is_done}
      onChange={async () => {
        if (task.is_done) {
          markTaskAsTodo(task.id);
        } else {
          markTaskAsDone(task.id);
        }
      }}
    />
  );
}
