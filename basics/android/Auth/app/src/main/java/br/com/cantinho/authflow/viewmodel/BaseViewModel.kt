package br.com.cantinho.authflow.viewmodel

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import kotlinx.coroutines.*

/**
 * Dispatchers.Main is set as the default CoroutineDispatcher for viewModelScope.
 *
 * Unit Testing viewModelScope
 * Dispatchers.Main uses the Android Looper.getMainLooper() method to run code in the UI thread.
 * That method is available in Instrumented Android tests but not in Unit tests.
 * Use the org.jetbrains.kotlinx:kotlinx-coroutines-test:$coroutines_version library
 * to replace the Coroutines Main Dispatcher by calling Dispatchers.setMain(dispatcher: CoroutineDispatcher)
 * with a TestCoroutineDispatcher that is available in
 * org.jetbrains.kotlinx:kotlinx-coroutines-test:$coroutines_version.
 * TestCoroutineDispatcher is a dispatcher that gives us control
 * of how coroutines are executed, being able to pause/resume execution and control its virtual clock.
 * It was added as an experimental API in Kotlin Coroutines v1.2.1.
 * You can read more about it in the documentation.
 * Don’t use Dispatchers.Unconfined as a replacement of Dispatchers.Main,
 * it will break all assumptions and timings for code that does use Dispatchers.Main.
 * Since a unit test should run well in isolation and without any side effects,
 * you should call Dispatchers.resetMain() and clean up the executor when the test finishes running.
 *
 * [https://medium.com/androiddevelopers/easy-coroutines-in-android-viewmodelscope-25bffb605471]
 *
 */
class BaseViewModel : ViewModel() {

    // To promote encapsulation and immutability, hide the MutableLiveData objects behind


    /**
     * Heavy operation that cannot be done in the Main Thread.
     */
    fun launchDataLoad() {
        viewModelScope.launch {
            sortList()
            // Modify UI
        }
    }

    suspend fun sortList() = withContext(Dispatchers.Default) {
        // Heavy work.
    }
}