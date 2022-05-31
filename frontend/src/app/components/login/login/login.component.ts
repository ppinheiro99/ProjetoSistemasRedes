import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TokenService } from "../../../services/token/token.service";
import { AuthService } from "../../../services/auth/auth.service";
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  validationForm: FormGroup;
  isLoggedIn = false;
  isLoginFailed = false;
  errorMessage = '';
  constructor(private router: Router, private authService: AuthService, private tokenStorage: TokenService, public fb: FormBuilder) {
    this.validationForm = fb.group({
      email: [null, [ Validators.required, Validators.email]],
      password: [null, Validators.required],
    })
  }


  get email() {
    return this.validationForm.get('email');
  }

  get password() {
    return this.validationForm.get('password');
  }

  ngOnInit() {
    if (this.tokenStorage.getToken()) {
      this.isLoggedIn = true;
      this.router.navigate(['/dashboard'])
    }
  }
  onLogin() {
    localStorage.setItem('isLoggedin', 'true');
    this.authService.login(this.validationForm.value).subscribe(
      data => {
        this.tokenStorage.saveToken(data.token);
        this.tokenStorage.saveUser(data);
        this.isLoginFailed = false;
        this.isLoggedIn = true;
        this.router.navigate(['/dashboard']);
      },
      err => {
        this.errorMessage = err.error.message;
        this.isLoginFailed = true;
      }
    );
  }
}
