<template>
  <div class="auth">
    <h1 class="auth__heading">
      Цей акаунт було видалено
    </h1>

    <p class="auth__description">
      Якщо ви вважаєте, що це помилка, будь ласка, зв'яжіться з підтримкою
    </p>
  </div>
</template>

<style scoped>
@layer page {
  .auth {
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
    gap: 12px;
  }

  .auth__heading {
    font-size: 28px;
  }

  .auth__description {
    max-width: 400px;
    text-align: center;
    text-wrap: balance;
  }
}
</style>
