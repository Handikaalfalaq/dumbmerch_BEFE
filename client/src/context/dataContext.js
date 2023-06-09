import React, { createContext, useState} from 'react';

export const DataContext = createContext();

// const initialState = {
//   isLogin: false,
//   user: {},
// };

// const reducer = (state, action) => {
//   const { type, payload } = action;
//   console.log("ini payload", payload)

//   switch (type) {
//     case "USER_SUCCESS":
//     case "LOGIN_SUCCESS":
//       localStorage.setItem("token", payload.token);
//       return {
//         isLogin: true,
//         user: payload,
//       };
//     case "AUTH_ERROR":
//     case "LOGOUT":
//       localStorage.removeItem("token");
//       return {
//         isLogin: false,
//         user: {},
//       };
//     default:
//       throw new Error();
//   }
// };

export const DataProvider = ({children}) => {
    // const [state, dispatch] = useReducer(reducer, initialState);
    const [dataUserLogin, setDataUserLogin] = useState([]);
    const [amount, setAmount] = useState('1');
    const [dateBooking, setDateBooking] = useState('Belum ada tanggal');
    const [formLogin, setFormLogin] = useState([]);
    const [userLogin, setUserLogin] = useState(false);
    const [adminLogin, setAdminLogin] = useState(false);
    const [navbarProfile, setNavbarProfile] = useState(false);
    const [Number, setNumber] = useState('');
    const [dataBooking, setDataBooking] = useState('');
    const [paySukses, setPaySukses] = useState(false);
    const [appearancePay, setAppearancePay] = useState(false);
    const [message, setMessage] = useState(null);
    const [showLoginModal, setShowLoginModal] = useState(false);
    const [idUserLogin, setIdUserLogin] = useState("");
    // console.log("ini state",state)

    return (
        <DataContext.Provider value={{ dataUserLogin, setDataUserLogin, amount, setAmount, dateBooking, setDateBooking, userLogin, setUserLogin, dataBooking, setDataBooking, paySukses, setPaySukses, adminLogin, setAdminLogin, navbarProfile, setNavbarProfile, Number, setNumber, appearancePay, setAppearancePay,  formLogin, setFormLogin, message, setMessage, showLoginModal, setShowLoginModal, idUserLogin, setIdUserLogin}}>
            {children}
        </DataContext.Provider>

    );
};
