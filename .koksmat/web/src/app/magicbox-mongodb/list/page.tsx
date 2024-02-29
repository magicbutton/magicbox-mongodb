"use client";
import { useProcess } from "@/koksmat/useprocess";
import { APPNAME } from "@/app/appid";

export default function Index() {
  const { isLoading, error, data } = useProcess(
    APPNAME,
    ["discover", "list"],
    20,
    "echo"
  );

  return (
    <div>
      <pre>{data}</pre>
    </div>
  );
}
