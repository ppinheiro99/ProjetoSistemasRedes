import { Component, OnInit, Inject } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { TokenService } from "../../services/token/token.service";
import { AuthService } from "../../services/auth/auth.service";
import { MatDialog, MAT_DIALOG_DATA } from '@angular/material/dialog';

export interface DialogData {
  passwordChangeData: boolean
}

@Component({
  selector: 'app-passrecover',
  templateUrl: './passrecover.component.html',
  styleUrls: ['./passrecover.component.scss']
})
export class PassrecoverComponent implements OnInit {

  validationForm: FormGroup;
  form: any = {};
  errorMessage = '';
  dialog: any;
  passwordChange = ''

  constructor(private router: Router, private authService: AuthService, private tokenStorage: TokenService, public fb: FormBuilder,public matDialog: MatDialog) {
    this.validationForm = fb.group({
      email: [null, [ Validators.required, Validators.email]]
    });
   }

  get email() {
    return this.validationForm.get('email');
  }
  onSubmit(): void {
    this.authService.passRecover(this.validationForm.value).subscribe(
      data => {
        this.passwordChange = 'true'
      },
      err => {
        this.passwordChange = 'false'
        this.errorMessage = err.error.message
      }
    );
  }
  ngOnInit(): void {
    this.passwordChange = ''
    
  }
  reloadPage(): void {
    window.location.reload();
  }

}

@Component({
  selector: 'dialog-error',
  templateUrl: 'dialog-error.html',
})
export class DialogError {
  constructor(@Inject(MAT_DIALOG_DATA) public data: DialogData) {
  }
}