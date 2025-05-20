import { useEffect, useState } from 'react'

export default function Home() {
  const [data, setData] = useState(null)

  useEffect(() => {
    fetch('/api/react')
      .then(res => res.json())
      .then(setData)
  }, [])

  return (
    <div>
      <h1>Armwrestling Guide</h1>
      <p>Welcome to the ultimate resource for armwrestling!</p>
      {data && <p><strong>API says:</strong> {data.message}</p>}
    </div>
  )
}