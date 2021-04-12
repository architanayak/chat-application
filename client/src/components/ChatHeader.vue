<template>
  <b-card title="Chat-App">
    <b-form>
      <label class="sr-only" for="inline-form-input-username">Username</label>
      <b-input-group prepend="@" class="mb-2 mr-sm-2 mb-sm-0">
        <b-form-input
          id="inline-form-input-username"
          placeholder="Username"
          v-model.trim="username"
        ></b-form-input>
        <b-button @click="joinRoom">Join</b-button>
      </b-input-group>
    </b-form>
  </b-card>
</template>

<style scoped></style>
