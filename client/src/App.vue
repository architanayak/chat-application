<template>
  <div id="app">
    <chat-navbar></chat-navbar>
    <br />
    <chat-chatbox :socket="socket"></chat-chatbox>
  </div>
</template>

<script>
import io from "socket.io-client";
import ChatNavbar from "./components/ChatNavbar.vue";
import ChatChatbox from "./components/ChatChatbox.vue";
export default {
  name: "App",
  components: {
    ChatNavbar,
    ChatChatbox,
  },
  data() {
    return {
      socket: null,
      serverUrl: process.env.VUE_APP_SOCKET_URL || "ws://localhost:8080",
    };
  },
  created() {
    this.connectToWebsocket();
  },
  methods: {
    connectToWebsocket() {
      this.socket = io(this.serverUrl, {
        transports: ["websocket"],
      });
    },
  },
};
</script>

<style>
html,
body,
#app,
.card {
  height: 100%;
}
</style>
