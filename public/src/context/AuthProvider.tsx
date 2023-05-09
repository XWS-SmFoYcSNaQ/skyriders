import { createContext, useState } from "react";


export interface AuthContextType {
    auth: any;
    setAuth: React.Dispatch<React.SetStateAction<any>>;
}

const AuthContext = createContext<AuthContextType>({
    auth: null,
    setAuth: () => { }
});

export const AuthProvider = ({ children }: any) => {
    const isAuthenticated = false;
    const [auth, setAuth] = useState({ isAuthenticated });

    return (
        <AuthContext.Provider value={{ auth, setAuth }}>
            {children}
        </AuthContext.Provider>
    )
}

export default AuthContext;