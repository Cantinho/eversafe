package br.com.cantinho.authflow

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import androidx.navigation.Navigation.findNavController

class StartActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_start)
    }

    /**
     * To ensure the Back button works properly, you also need to override the [onSupportNavigateUp] method in
     * [StartActivity].
     */
    override fun onSupportNavigateUp() =
        findNavController(this, R.id.start_nav_host_fragment).navigateUp()
}
