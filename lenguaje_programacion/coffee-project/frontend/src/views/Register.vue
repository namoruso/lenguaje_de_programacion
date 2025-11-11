<template>
  <div class="auth-container">
    <div class="coffee-beans"></div>
    <div class="auth-card">
      <div class="logo">
        <div class="coffee-cup">☕</div>
        <h1>Coffee House</h1>
        <p class="tagline">Tu momento perfecto</p>
      </div>
      
      <form @submit.prevent="handleLogin">
        <div class="form-group">
          <label>📧 Email</label>
          <input 
            v-model="email" 
            type="email" 
            placeholder="tu@email.com"
            required
          />
        </div>

        <div class="form-group">
          <label>🔒 Contraseña</label>
          <input 
            v-model="password" 
            type="password" 
            placeholder="••••••••"
            required
          />
        </div>

        <button type="submit" class="btn-primary">
          Iniciar Sesión
        </button>
        
        <p v-if="errorMessage" class="error">{{ errorMessage }}</p>
      </form>

      <div class="divider">
        <span>o</span>
      </div>

      <p class="switch-auth">
        ¿Nuevo aquí? 
        <router-link to="/register">Crear cuenta</router-link>
      </p>
    </div>
    
    <div class="footer-quote">
      "La vida es demasiado corta para tomar mal café"
    </div>
  </div>
</template>

<script>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import api from '../services/api';

export default {
  name: 'Login',
  setup() {
    const router = useRouter();
    const email = ref('');
    const password = ref('');
    const errorMessage = ref('');

    const handleLogin = async () => {
      try {
        errorMessage.value = '';
        
        const response = await api.post('/auth/login', {
          email: email.value,
          password: password.value
        });

        if (response.data.success) {
          localStorage.setItem('user', JSON.stringify(response.data.user));
          router.push('/menu');
        }
      } catch (error) {
        errorMessage.value = error.response?.data?.message || 'Error al iniciar sesión';
      }
    };

    return {
      email,
      password,
      errorMessage,
      handleLogin
    };
  }
};
</script>

<style scoped>
.auth-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #6B4423 0%, #3E2723 100%);
  position: relative;
  overflow: hidden;
  padding: 20px;
}

.coffee-beans {
  position: absolute;
  width: 100%;
  height: 100%;
  opacity: 0.05;
  background-image: 
    radial-gradient(circle at 20% 30%, #D4A574 0%, transparent 50%),
    radial-gradient(circle at 80% 70%, #D4A574 0%, transparent 50%),
    radial-gradient(circle at 50% 50%, #D4A574 0%, transparent 50%);
  pointer-events: none;
}

.auth-card {
  background: #FFF8E7;
  padding: 50px 40px;
  border-radius: 20px;
  box-shadow: 
    0 20px 60px rgba(0,0,0,0.4),
    inset 0 1px 0 rgba(255,255,255,0.6);
  width: 100%;
  max-width: 420px;
  position: relative;
  border: 3px solid #D4A574;
}

.logo {
  text-align: center;
  margin-bottom: 35px;
}

.coffee-cup {
  font-size: 60px;
  margin-bottom: 10px;
  animation: steam 2s ease-in-out infinite;
  display: inline-block;
}

@keyframes steam {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(-5px); }
}

.logo h1 {
  color: #3E2723;
  margin: 0;
  font-family: 'Georgia', serif;
  font-size: 32px;
  font-weight: bold;
  text-shadow: 2px 2px 4px rgba(0,0,0,0.1);
}

.tagline {
  color: #6B4423;
  font-style: italic;
  margin-top: 5px;
  font-size: 14px;
}

.form-group {
  margin-bottom: 25px;
}

label {
  display: block;
  margin-bottom: 8px;
  color: #3E2723;
  font-weight: 600;
  font-size: 14px;
}

input {
  width: 100%;
  padding: 14px 16px;
  border: 2px solid #D4A574;
  border-radius: 10px;
  font-size: 16px;
  background: white;
  transition: all 0.3s;
  color: #3E2723;
}

input:focus {
  outline: none;
  border-color: #8B5A3C;
  box-shadow: 0 0 0 3px rgba(139, 90, 60, 0.1);
  transform: translateY(-2px);
}

input::placeholder {
  color: #A78B71;
}

.btn-primary {
  width: 100%;
  padding: 16px;
  background: linear-gradient(135deg, #8B5A3C 0%, #6B4423 100%);
  color: #FFF8E7;
  border: none;
  border-radius: 10px;
  font-size: 18px;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.3s;
  box-shadow: 0 4px 15px rgba(107, 68, 35, 0.4);
  text-transform: uppercase;
  letter-spacing: 1px;
}

.btn-primary:hover {
  background: linear-gradient(135deg, #6B4423 0%, #3E2723 100%);
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(107, 68, 35, 0.6);
}

.btn-primary:active {
  transform: translateY(0px);
}

.error {
  color: #C62828;
  text-align: center;
  margin-top: 15px;
  padding: 10px;
  background: #FFEBEE;
  border-radius: 8px;
  font-size: 14px;
  border-left: 4px solid #C62828;
}

.divider {
  text-align: center;
  margin: 30px 0;
  position: relative;
}

.divider::before,
.divider::after {
  content: '';
  position: absolute;
  top: 50%;
  width: 45%;
  height: 1px;
  background: #D4A574;
}

.divider::before {
  left: 0;
}

.divider::after {
  right: 0;
}

.divider span {
  background: #FFF8E7;
  padding: 0 15px;
  color: #6B4423;
  font-weight: 600;
  position: relative;
}

.switch-auth {
  text-align: center;
  margin-top: 25px;
  color: #6B4423;
  font-size: 15px;
}

.switch-auth a {
  color: #8B5A3C;
  text-decoration: none;
  font-weight: 700;
  transition: color 0.3s;
  border-bottom: 2px solid transparent;
}

.switch-auth a:hover {
  color: #3E2723;
  border-bottom: 2px solid #8B5A3C;
}

.footer-quote {
  position: absolute;
  bottom: 30px;
  left: 50%;
  transform: translateX(-50%);
  color: #D4A574;
  font-style: italic;
  font-size: 14px;
  text-align: center;
  max-width: 400px;
  text-shadow: 2px 2px 4px rgba(0,0,0,0.5);
}

@media (max-width: 480px) {
  .auth-card {
    padding: 40px 25px;
  }
  
  .logo h1 {
    font-size: 28px;
  }
  
  .footer-quote {
    bottom: 15px;
    font-size: 12px;
    padding: 0 20px;
  }
}
</style>