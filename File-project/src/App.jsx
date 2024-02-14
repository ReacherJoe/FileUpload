import React, { useState , useEffect} from 'react';
import axios from 'axios';


function FileUpload() {
    const [uploadHistory, setUploadHistory] = useState([]);

    useEffect (() => {
      axios.get('http://localhost:8080/data')
    
      .then(response => {
        console.log(response.data)
        setUploadHistory(response.data)
      })
      .catch(error => console.error('Error fetching data:', error));
    }, []);
    
    const handleSubmit = async (e) => {
        e.preventDefault();
        const formData = new FormData(e.target);
        try {
            const response = await axios.post("http://localhost:8080/upload", formData);
            console.log(response.data);
            
            setUploadHistory(prevUploadHistory => [...prevUploadHistory, response.data]);
        } catch (error) {
            console.log(error);
        }
    };

    return (
      <>
        <div>
            <h1>File upload</h1>
            <form onSubmit={handleSubmit}>
                <input type="file" name="file" /> <br/><br/>
                <button type="submit">Upload</button>
            </form>

            <h2>Upload History</h2>
            <ul>
                {uploadHistory.map((data, index) => (
                      <li key={index}>
                     <a href={data.images} target="_blank" rel="noopener noreferrer">
                {data.images}
                    </a>
                      </li>
                   
                ))}
            </ul>
        </div>
        </>
    );
}

export default FileUpload;
