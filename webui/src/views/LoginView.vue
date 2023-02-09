<template>
	<div
		class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
		<h1 class="h2">Login with your username!</h1>
		<div class="btn-toolbar mb-2 mb-md-0">
			<div class="btn-group me-2">
<!--				<button type="button" class="btn btn-sm btn-outline-secondary" @click="refresh">-->
<!--					Refresh-->
<!--				</button>-->
				<button type="button" class="btn btn-sm btn-outline-secondary" @click="login">
					Sign in
				</button>
			</div>
		</div>
	</div>

	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

</template>

<script>
export default {
	name: "LoginView",
	data: function () {
		return {
			errormsg: null,
			user_name: "",
			user_profile: {
				user_id: 0,
				user_name: "username...",
				pass_word: "not required",
			},
		}
	},
	methods: {
		async login() {
			// Perform login.
			if (this.user_name == "") {
				this.errormsg = "Please input a valid username";
			} else {
				try {
					let res = await this.$axios.post("/session", {user_name: this.user_name})

					this.user_profile = res.data
					localStorage.setItem("user_name", this.user_profile.user_name);

					// Move to main screen after authentication, i.e., the photo stream
					this.$router.push({path: '/users/:user_id/stream'})
				} catch (error) {
					if (error.response && error.response.status === 400)
						this.errormsg = "Bad Request. Try again.";
					if (error.response && error.response.status === 404) {
						this.errormsg = "Username Not Found. Try again.";
					} else if (error.response && error.response.status === 500) {
						this.errormsg = "Internal Server Error. Try again at a later time.";
					}
				}
			}
		}
	},
}
</script>

<style scoped>

</style>
