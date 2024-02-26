
/*
This file is in testing stage. Still figuring out how to integrate node with goLang
Basically what has been done here is we have two servers one for node running on port 8000 
other of go running on 3000. Only the routes on node are exposed for requests and the incoming request
is then forwared to the respective request routes in go. Only the internal code knows where the request is going
This has been purely my understanding of the assignemtn and the real implementation may differ significantly
*/

import express from 'express';
const app = express();
const PORT = 8000;
import fetch from 'node-fetch';
import bodyParser from 'body-parser';


// Parse incoming request bodies in a middleware before  handlers
//this is done so we direcrly access req.body

app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());

app.get('/',(req, res)=>{
  res.status(200).json("hello")
})
const bankRouter = express.Router()


bankRouter.get("/", async(req,res)=>{
  console.log("request body is ", req.body)
  try {
    const response = await fetch('http://localhost:3000/bank', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(req.body)
    });
    
    if (response.ok) {
      const data = await response.json();
      res.status(200).json({ message: 'Successfully called Go API', data });
    } else {
      throw new Error('Failed to fetch data from Go API');
    }
  } catch (error) {
    res.status(500).json({ message: 'Failed to call Go API', error: error.message });
  }
});

bankRouter.post('/', async (req, res) => {
  console.log("request body is ", req.body)
  try {
    const response = await fetch('http://localhost:3000/bank', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(req.body)
    });
    
    if (response.ok) {
      const data = await response.json();
      res.status(200).json({ message: 'Successfully called Go API', data });
    } else {
      throw new Error('Failed to fetch data from Go API');
    }
  } catch (error) {
    res.status(500).json({ message: 'Failed to call Go API', error: error.message });
  }
});



app.use("/bank",bankRouter)
app.listen(PORT, () => {
  console.log(`Server running at: http://localhost:${PORT}/`);
});
