// import Greetings from "./Greetings";

// export default function App(){
//   return (
//     <>
//       <h1>
//         <marquee >Welcome to REACT
//         </marquee>
//       </h1>
//       <Greetings />
//     </>
//   )
// }


import {BrowserRouter,Routes,Route} from "react-router-dom";
import CarList from "./cars/CarList";
import CarCreate from "./cars/CarCreate";
import CarView from "./cars/CarView";

function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
            <Route path="" element={<CarList/>}/>
            <Route path="/car/list" element={<CarList/>}/>
            <Route path="/car/create" element={<CarCreate/>}/>
            <Route path="/car/view" element={<CarView/>}/>
        </Routes>
      
      </BrowserRouter>
    </>
  )
}

export default App