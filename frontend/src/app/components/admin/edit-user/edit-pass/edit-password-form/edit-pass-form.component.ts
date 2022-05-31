import { Component, OnInit} from '@angular/core';
import { AuthService } from "../../../../../services/auth/auth.service";
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import { MatDialog } from '@angular/material/dialog';
import { TokenService } from "../../../../../services/token/token.service";

@Component({
  selector: 'app-edit-pass-form',
  templateUrl: './edit-pass-form.component.html',
  styleUrls: ['./edit-pass-form.component.scss']
})
export class EditPassFormComponent implements OnInit {
  validationForm: FormGroup
  form: any = {}
  isSuccessful = false
  passwordFailed = false
  newpasswordFailed = false
  newpasswordConfFailed = false
  missmatchpass = false
  samepassword = false
  isSignUpFailed = false
  errorMessage = '';
  check_email = '';
  check_pass = '';
  check_pass2 = '';
  check_pass3 = '';
  user;
 
  constructor(private tokenService: TokenService ,private authService: AuthService, public fb: FormBuilder, public dialog: MatDialog) {
    this.validationForm = fb.group({
      password_old: [null, Validators.required],
      password_new: [null, Validators.required],
      password_newConf: [null, Validators.required],
    });
  }
  
  get oldpassword() { return this.validationForm.get('password_old') }
  get newpassword() { return this.validationForm.get('password_new') }
  get newpasswordC() { return this.validationForm.get('password_newConf') }

  ngOnInit(): void {
  }

  onSubmit(): void {
    
    this.passwordFailed = false
    this.newpasswordFailed = false
    this.isSignUpFailed = false
    
    this.user = this.tokenService.getUser()
    
    this.authService.changePassword(this.validationForm.value,this.user.Email).subscribe(
      data => {
        this.isSuccessful = true
        this.isSignUpFailed = false
        this.reloadPage()
      },
      err => {
        this.errorMessage = err.error.message
        this.isSignUpFailed = true
        if(err.error.message == "Invalid password!"){
         
          this.passwordFailed = true
        }else if(err.error.message == "Password dont match!"){
          this.missmatchpass = true
        }
        else{
          this.samepassword = true
        }
      }
    );
  }

  reloadPage(): void {
    window.location.reload();
  }


}
