package br.com.cantinho.authflow

import android.content.Intent
import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.util.Log
import com.amazonaws.mobile.client.*

class AuthenticationActivity : AppCompatActivity() {

    private val TAG = AuthenticationActivity::class.java.simpleName


    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_authentication)

        AWSMobileClient.getInstance().initialize(applicationContext, object : Callback<UserStateDetails> {

            override fun onResult(userStateDetails: UserStateDetails?) {
                Log.i(TAG, userStateDetails?.userState.toString())
                when (userStateDetails?.userState) {
                    UserState.GUEST -> {
                        Log.i("userState", "user is in guest mode")
                    }
                    UserState.SIGNED_IN -> {
                        Log.i("userState", "user is signed in")
                        val i = Intent(this@AuthenticationActivity, MainActivity::class.java)
                        startActivity(i)
                    }
                    UserState.SIGNED_OUT -> {
                        Log.i("userState", "user is signed out")
                        showSignIn()
                    }
                    UserState.SIGNED_OUT_USER_POOLS_TOKENS_INVALID -> {
                        Log.i("userState", "need to login again")
                        AWSMobileClient.getInstance().signOut()
                        showSignIn()
                    }
                    UserState.SIGNED_OUT_FEDERATED_TOKENS_INVALID -> {
                        Log.i("userState", "user logged in via federation, but currently needs new tokens")
                        AWSMobileClient.getInstance().signOut()
                        showSignIn()
                    }
                    else -> {
                        Log.e("userState", "unsupported")
                        AWSMobileClient.getInstance().signOut()
                        showSignIn()
                    }
                }
            }

            override fun onError(e: java.lang.Exception?) {
                Log.e(TAG, e.toString())
            }


        })
    }

    private fun showSignIn() {
        try {
            AWSMobileClient.getInstance().showSignIn(
                this,
                SignInUIOptions.builder().nextActivity(MainActivity::class.java).build()
            )
        } catch (e: Exception) {
            Log.e(TAG, e.toString())
        }

    }
}
