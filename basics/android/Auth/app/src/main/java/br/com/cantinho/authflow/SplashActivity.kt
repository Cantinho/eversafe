package br.com.cantinho.authflow

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.os.Handler
import android.util.Log
import android.view.View
import kotlinx.android.synthetic.main.activity_splash.*


/**
 * An example full-screen activity that shows and hides the system UI (i.e.
 * status bar and navigation/system bar) with user interaction.
 */
class SplashActivity : AppCompatActivity() {
    private val mHideHandler = Handler()

    private var mVisible: Boolean = false
    private val mHideRunnable = Runnable { hide() }
    private val mAnimateRunnable = Runnable { animate() }


    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        setContentView(R.layout.activity_splash)
        supportActionBar?.setDisplayHomeAsUpEnabled(true)

        mVisible = true
    }

    override fun onPostCreate(savedInstanceState: Bundle?) {
        super.onPostCreate(savedInstanceState)

        // Trigger the initial hide() shortly after the activity has been
        // created, to briefly hint to the user that UI controls
        // are available.
        delayedInit(UI_ANIMATION_HIDE_DELAY, UI_ANIMATION_DELAY)
    }

    private fun animate() {
        Log.i("animate", " start motion layout")
        splashscreen.transitionToEnd()
    }

    private fun hide() {
        // Hide UI first
        supportActionBar?.hide()
        mVisible = false

        // Remove the status and navigation bar
        // Note that some of these constants are new as of API 16 (Jelly Bean)
        // and API 19 (KitKat). It is safe to use them, as they are inlined
        // at compile-time and do nothing on earlier devices.
        fullscreen_content.systemUiVisibility =
            View.SYSTEM_UI_FLAG_LOW_PROFILE or
                    View.SYSTEM_UI_FLAG_FULLSCREEN or
                    View.SYSTEM_UI_FLAG_LAYOUT_STABLE or
                    View.SYSTEM_UI_FLAG_IMMERSIVE_STICKY or
                    View.SYSTEM_UI_FLAG_LAYOUT_HIDE_NAVIGATION or
                    View.SYSTEM_UI_FLAG_HIDE_NAVIGATION
    }

    /**
     * Schedules a call to hide() in [delayedHideBarMillis] and animate() in [delayedAnimationMillis], canceling any
     * previously scheduled calls.
     */
    private fun delayedInit(delayedHideBarMillis: Int, delayedAnimationMillis: Int) {
        mHideHandler.removeCallbacks(mHideRunnable)
        mHideHandler.removeCallbacks(mAnimateRunnable)
        mHideHandler.postDelayed(mHideRunnable, delayedHideBarMillis.toLong())
        mHideHandler.postDelayed(mAnimateRunnable, delayedAnimationMillis.toLong())
    }

    companion object {

        /**
         * Some older devices needs a small delay between UI widget updates
         * and a change of the status and navigation bar.
         * [UI_ANIMATION_HIDE_DELAY] must be zero.
         */
        private val UI_ANIMATION_HIDE_DELAY = 0

        private val UI_ANIMATION_DELAY = 1000
    }
}
