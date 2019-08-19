package br.com.cantinho.authflow

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.fragment.app.Fragment
import androidx.navigation.findNavController
import kotlinx.android.synthetic.main.fragment_authentication.*


class AuthenticationFragment : Fragment() {

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        // Inflate the layout for this fragment
        return inflater.inflate(R.layout.fragment_authentication, container, false)
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        btnGoToHome.setOnClickListener { view -> view.findNavController().navigate(R.id.actionToHome) }
        btnGoToSignUp.setOnClickListener { view -> view.findNavController().navigate(R.id.actionToSignUp) }
        btnGoToForgotPassword.setOnClickListener { view -> view.findNavController().navigate(R.id.actionToForgotPassword) }
    }
}
