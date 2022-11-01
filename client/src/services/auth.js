import axios from 'axios'

class AuthServices {
    register(){

    }
    async login(user){
        return await axios.post("https://universityweb.site/api/login", {email: user.email, password: user.password}).then(resp => {
            console.log('OKAY');
        }).catch(err => {
            console.log(err.response.data.msg);
        })
    }
    logout(){

    } 
}

export default new AuthServices()