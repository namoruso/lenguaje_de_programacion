<template>
  <div class="menu-container">
    <header class="menu-header">
      <div class="header-content">
        <div class="brand">
          <span class="coffee-icon">☕</span>
          <div>
            <h1>Coffee House</h1>
            <p class="subtitle">Menú de especialidades</p>
          </div>
        </div>
        
        <div class="user-info">
          <div class="welcome">
            <span class="user-icon">👋</span>
            <span class="username">{{ user?.username }}</span>
          </div>
          <button @click="handleLogout" class="btn-logout">
            🚪 Salir
          </button>
        </div>
      </div>
    </header>

    <div class="menu-intro">
      <h2>Nuestras especialidades</h2>
      <p>Cada taza es una experiencia única preparada con amor</p>
    </div>

    <div v-if="loading" class="loading">
      <div class="coffee-loader">☕</div>
      <p>Preparando tu menú...</p>
    </div>

    <div v-else class="coffee-grid">
      <div 
        v-for="coffee in coffees" 
        :key="coffee.id" 
        class="coffee-card"
      >
        <div class="coffee-badge">⭐ Popular</div>
        
        <div class="coffee-image-container">
          <img 
            :src="coffee.image" 
            :alt="coffee.name"
            class="coffee-image"
            @error="handleImageError"
          />
          <div class="image-overlay">
            <span class="view-detail">Ver detalles</span>
          </div>
        </div>
        
        <div class="coffee-content">
          <h3>{{ coffee.name }}</h3>
          <p class="description">{{ coffee.description }}</p>
          
          <div class="rating">
            <span class="stars">★★★★★</span>
            <span class="reviews">(4.8)</span>
          </div>
          
          <div class="card-footer">
            <p class="price">
              <span class="currency">$</span>
              <span class="amount">{{ coffee.price.toFixed(2) }}</span>
            </p>
            <button class="btn-order">
              🛒 Ordenar
            </button>
          </div>
        </div>
      </div>
    </div>

    <footer class="page-footer">
      <p>☕ Coffee House © 2025 - El arte del buen café</p>
    </footer>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import api from '../services/api';

export default {
  name: 'CoffeeMenu',
  setup() {
    const router = useRouter();
    const coffees = ref([]);
    const loading = ref(true);
    const user = ref(null);

    onMounted(async () => {
      const userData = localStorage.getItem('user');
      if (userData) {
        user.value = JSON.parse(userData);
      }

      try {
        const response = await api.get('/coffees');
        if (response.data.success) {
          coffees.value = response.data.coffees;
        }
      } catch (error) {
        console.error('Error al cargar cafés:', error);
      } finally {
        loading.value = false;
      }
    });

    const handleLogout = () => {
      localStorage.removeItem('user');
      router.push('/login');
    };

    const handleImageError = (event) => {
      event.target.src = 'https://via.placeholder.com/400x300/8B5A3C/FFF8E7?text=Coffee';
    };

    return {
      coffees,
      loading,
      user,
      handleLogout,
      handleImageError
    };
  }
};
</script>

<style scoped>
.menu-container {
  min-height: 100vh;
  background: linear-gradient(to bottom, #FFF8E7 0%, #F5E6D3 100%);
}

.menu-header {
  background: linear-gradient(135deg, #3E2723 0%, #6B4423 100%);
  color: #FFF8E7;
  padding: 25px 30px;
  box-shadow: 0 4px 20px rgba(0,0,0,0.3);
  position: sticky;
  top: 0;
  z-index: 100;
  border-bottom: 3px solid #D4A574;
}

.header-content {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 20px;
}

.brand {
  display: flex;
  align-items: center;
  gap: 15px;
}

.coffee-icon {
  font-size: 45px;
  animation: rotate 10s linear infinite;
}

@keyframes rotate {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.brand h1 {
  margin: 0;
  font-family: 'Georgia', serif;
  font-size: 28px;
  text-shadow: 2px 2px 4px rgba(0,0,0,0.3);
}

.subtitle {
  margin: 0;
  font-size: 13px;
  opacity: 0.9;
  font-style: italic;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 20px;
}

.welcome {
  display: flex;
  align-items: center;
  gap: 8px;
  background: rgba(255, 248, 231, 0.15);
  padding: 8px 16px;
  border-radius: 25px;
  border: 1px solid rgba(212, 165, 116, 0.3);
}

.user-icon {
  font-size: 20px;
}

.username {
  font-weight: 600;
  color: #D4A574;
}

.btn-logout {
  padding: 10px 20px;
  background: rgba(212, 165, 116, 0.2);
  color: #FFF8E7;
  border: 2px solid #D4A574;
  border-radius: 10px;
  cursor: pointer;
  font-weight: 600;
  transition: all 0.3s;
  font-size: 14px;
}

.btn-logout:hover {
  background: #D4A574;
  color: #3E2723;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(212, 165, 116, 0.4);
}

.menu-intro {
  text-align: center;
  padding: 50px 20px 30px;
  max-width: 800px;
  margin: 0 auto;
}

.menu-intro h2 {
  color: #3E2723;
  font-size: 36px;
  margin-bottom: 10px;
  font-family: 'Georgia', serif;
}

.menu-intro p {
  color: #6B4423;
  font-size: 18px;
  font-style: italic;
}

.loading {
  text-align: center;
  padding: 80px 20px;
}

.coffee-loader {
  font-size: 60px;
  animation: bounce 1s ease-in-out infinite;
  display: inline-block;
}

@keyframes bounce {
  0%, 100% { transform: translateY(0); }
  50% { transform: translateY(-20px); }
}

.loading p {
  margin-top: 20px;
  color: #6B4423;
  font-size: 18px;
  font-style: italic;
}

.coffee-grid {
  max-width: 1300px;
  margin: 0 auto;
  padding: 30px 20px 80px;
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 35px;
}

.coffee-card {
  background: white;
  border-radius: 20px;
  box-shadow: 0 8px 25px rgba(107, 68, 35, 0.15);
  overflow: hidden;
  display: flex;
  flex-direction: column;
  height: 500px;
  transition: all 0.4s ease;
  border: 2px solid #F5E6D3;
  position: relative;
}

.coffee-card:hover {
  transform: translateY(-10px);
  box-shadow: 0 15px 40px rgba(107, 68, 35, 0.25);
  border-color: #D4A574;
}

.coffee-badge {
  position: absolute;
  top: 15px;
  right: 15px;
  background: linear-gradient(135deg, #FFD700 0%, #FFA500 100%);
  color: #3E2723;
  padding: 6px 12px;
  border-radius: 20px;
  font-size: 12px;
  font-weight: bold;
  z-index: 10;
  box-shadow: 0 2px 8px rgba(0,0,0,0.2);
}

.coffee-image-container {
  width: 100%;
  height: 220px;
  overflow: hidden;
  background: linear-gradient(135deg, #F5E6D3 0%, #D4A574 100%);
  position: relative;
}

.coffee-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
  transition: transform 0.4s ease;
}

.coffee-card:hover .coffee-image {
  transform: scale(1.1);
}

.image-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(62, 39, 35, 0.7);
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0;
  transition: opacity 0.3s;
}

.coffee-card:hover .image-overlay {
  opacity: 1;
}

.view-detail {
  color: white;
  font-weight: bold;
  font-size: 16px;
  padding: 12px 24px;
  border: 2px solid white;
  border-radius: 25px;
  background: rgba(212, 165, 116, 0.3);
}

.coffee-content {
  padding: 25px;
  display: flex;
  flex-direction: column;
  flex-grow: 1;
  background: linear-gradient(to bottom, white 0%, #FFF8E7 100%);
}

.coffee-content h3 {
  color: #3E2723;
  margin-bottom: 12px;
  font-size: 24px;
  min-height: 32px;
  font-family: 'Georgia', serif;
}

.description {
  color: #6B4423;
  margin-bottom: 15px;
  line-height: 1.6;
  flex-grow: 1;
  font-size: 14px;
}

.rating {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 15px;
}

.stars {
  color: #FFD700;
  font-size: 16px;
  letter-spacing: 2px;
}

.reviews {
  color: #8B5A3C;
  font-size: 13px;
  font-weight: 600;
}

.card-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 15px;
  border-top: 2px dashed #D4A574;
  margin-top: auto;
}

.price {
  margin: 0;
  color: #3E2723;
  font-weight: bold;
  display: flex;
  align-items: baseline;
  gap: 2px;
}

.currency {
  font-size: 20px;
}

.amount {
  font-size: 32px;
}

.btn-order {
  padding: 12px 24px;
  background: linear-gradient(135deg, #8B5A3C 0%, #6B4423 100%);
  color: #FFF8E7;
  border: none;
  border-radius: 25px;
  font-size: 14px;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.3s;
  box-shadow: 0 4px 12px rgba(107, 68, 35, 0.3);
}

.btn-order:hover {
  background: linear-gradient(135deg, #6B4423 0%, #3E2723 100%);
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(107, 68, 35, 0.4);
}

.page-footer {
  background: #3E2723;
  color: #D4A574;
  text-align: center;
  padding: 30px;
  border-top: 3px solid #D4A574;
}

.page-footer p {
  margin: 0;
  font-style: italic;
  font-size: 14px;
}

@media (max-width: 768px) {
  .header-content {
    flex-direction: column;
    text-align: center;
  }
  
  .brand {
    flex-direction: column;
    gap: 10px;
  }
  
  .user-info {
    flex-direction: column;
    width: 100%;
  }
  
  .welcome {
    width: 100%;
    justify-content: center;
  }
  
  .btn-logout {
    width: 100%;
  }
  
  .coffee-grid {
    grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
    gap: 25px;
    padding: 20px;
  }
  
  .menu-intro h2 {
    font-size: 28px;
  }
  
  .menu-intro p {
    font-size: 16px;
  }
}
</style>