"use server";

import { refresh } from "next/cache";

export async function markTaskAsDone(taskID: string) {
  // @TODO: handle not ok status response
  await fetch(`http://localhost:8080/api/tasks/${taskID}/status/done`, {
    method: "PATCH",
  });

  refresh();
}

export async function markTaskAsTodo(taskID: string) {
  // @TODO: handle not ok status response
  await fetch(`http://localhost:8080/api/tasks/${taskID}/status/todo`, {
    method: "PATCH",
  });

  refresh();
}
