"use client"

import { useContext, useEffect, useState } from "react"
import { Context } from "../context"
import { Button } from "@/components/ui/button"
import { set } from "date-fns"
import { useProcess } from "@/koksmat/useprocess"
import { APPNAME } from "@/app/appid"


export interface Blob{
  size: number
  lastModified: string
  name: string
}

export default function Index() {
  const context = useContext(Context)

  const { isLoading, error, data } = useProcess(
    APPNAME,
    ["restore","list"],    
    20,
    "echo"
  )
  const [items, setitems] = useState<Blob[]>([])

  useEffect(() => {
    if (!data)  return
    try {
      setitems(JSON.parse(data))
    } 
    catch (e) {
      console.error(e)
    }
    
  
   
  }, [data])
  
  return (
    <div>
      {error && <div className="text-red-600">{error}</div>}

      <pre>{data}</pre>
     {items.map((item, index) => (
       <div key={index}>
          <p>{item.name}</p>
          <p>{item.lastModified}</p>
          <p>{item.size}</p>  
          <Button onClick={() => alert("Not implemented")}>Restore</Button>
       

    </div>
     ))}
  </div>)
}
