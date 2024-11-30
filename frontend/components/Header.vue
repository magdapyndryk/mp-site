<template>
  <header class="header">
    <div class="logo">
      <img src="/logo.png" alt="Logo" />
    </div>
    <nav class="nav-menu">
      <ul>
        <li><NuxtLink to="/">O mnie</NuxtLink></li>
        <li><NuxtLink to="/programming">Programowanie</NuxtLink></li>
        <li><NuxtLink to="/motoring">Zmotoryzowanie</NuxtLink></li>
        <li><NuxtLink to="/personal-development">Rozwój osobisty</NuxtLink></li>
        <li><NuxtLink to="/my-brands">Moje marki</NuxtLink></li>
        <li><NuxtLink to="/shop">Sklep</NuxtLink></li>
        <li><NuxtLink to="/contact">Kontakt</NuxtLink></li>
      </ul>
    </nav>
  </header>
</template>

<style scoped>
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
  background-color: #f8f8f8;
}
.logo img {
  height: 50px;
}
.nav-menu ul {
  display: flex;
  list-style: none;
  gap: 1rem;
}
.nav-menu li {
  font-weight: bold;
}
</style>
