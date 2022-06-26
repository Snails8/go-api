import React, { useEffect, useState } from 'react';

export const Dashboard: React.FC = () => {
    const [message, setMessage] = useState("")
  
    useEffect(() => {
      fetch('/api/v1/', {
        method: 'GET', 
      })
      .then(res => res.json() )
      .then(data => {setMessage(data)})
    }, [])
  
    console.log(message)

    return (
        <>
            <div>aaa</div>        
        </>
    )
}