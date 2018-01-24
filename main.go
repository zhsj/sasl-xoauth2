package main

/*
#cgo CFLAGS: -Wall

#include <sasl/sasl.h>
#include <sasl/saslplug.h>

extern void getAuthString(char *, char *, char **, unsigned *);

static int get_cb_val(const sasl_utils_t *utils, unsigned int id, const char **result, unsigned *result_len) {
	int ret = SASL_FAIL;
	switch(id) {
	case SASL_CB_AUTHNAME:
		{
			sasl_getsimple_t *simple_cb;
			void *simple_context;
			ret = utils->getcallback(utils->conn, id, (sasl_callback_ft *)&simple_cb, &simple_context);
			ret = simple_cb(simple_context, id, result, result_len);
		}
		break;
	case SASL_CB_PASS:
		{
			sasl_getsecret_t *pass_cb;
			void *pass_context;
			sasl_secret_t *password;
			ret = utils->getcallback(utils->conn, SASL_CB_PASS, (sasl_callback_ft *)&pass_cb, &pass_context);
			ret = pass_cb(utils->conn, pass_context, SASL_CB_PASS, &password);
			*result = (char *)password->data;
			*result_len = password->len;
		}
		break;
	}
	return ret;
}


static int client_mech_new(void *glob_context, sasl_client_params_t *params, void **conn_context) {
	return SASL_OK;
}

static int client_mech_step(void *conn_context, sasl_client_params_t *params,
		const char *serverin, unsigned serverinlen, sasl_interact_t **prompt_need,
		const char **clientout, unsigned *clientoutlen, sasl_out_params_t *oparams) {
	int ret = SASL_FAIL;
	char * authname;
	unsigned authname_len;
	char * refresh_token;
	unsigned refresh_token_len;
	ret = get_cb_val(params->utils, SASL_CB_AUTHNAME, (const char **)&authname, &authname_len);
	ret = get_cb_val(params->utils, SASL_CB_PASS, (const char **)&refresh_token, &refresh_token_len);
	ret = params->canon_user(params->utils->conn, authname, authname_len, SASL_CU_AUTHID | SASL_CU_AUTHZID, oparams);
	char * auth;
	unsigned auth_len;
	getAuthString(authname, refresh_token, &auth, &auth_len);
	*clientout = auth;
	*clientoutlen = auth_len;
	return ret;
}

static void client_mech_dispose(void *conn_context, const sasl_utils_t *utils) {
}

static sasl_client_plug_t client_plugins[] =
{
	{
	"XOAUTH2",
	0,
	SASL_SEC_MAXIMUM,
	SASL_FEAT_WANT_CLIENT_FIRST | SASL_FEAT_ALLOWS_PROXY,
	NULL,
	NULL,
	&client_mech_new,
	&client_mech_step,
	&client_mech_dispose,
	NULL,
	NULL,
	NULL,
	NULL
	}
};

static inline int client_plug_init(const sasl_utils_t *utils, int maxversion, int *out_version, sasl_client_plug_t **pluglist, int *plugcount) {
	if (maxversion < SASL_CLIENT_PLUG_VERSION) {
		utils->seterror(utils->conn, 0, "XOAUTH2: Version mismatch");
		return SASL_BADVERS;
	}
	*out_version = SASL_CLIENT_PLUG_VERSION;
	*pluglist = client_plugins;
	*plugcount = 1;
	return SASL_OK;
}
*/
import "C"

//export sasl_client_plug_init
func sasl_client_plug_init(utils *C.sasl_utils_t, maxversion C.int, out_version *C.int, pluglist **C.sasl_client_plug_t, plugcount *C.int) C.int {
	//test()
	//printToken()
	return C.client_plug_init(utils, maxversion, out_version, pluglist, plugcount)
}

//export getAuthString
func getAuthString(user *C.char, refreshToken *C.char, auth **C.char, authLen *C.uint) {
	authString := genAuthString(C.GoString(user), C.GoString(refreshToken))
	*authLen = C.uint(len(authString))
	*auth = C.CString(authString)
}

func main() {
}
