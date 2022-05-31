import { Component, OnInit} from '@angular/core';
import { AuthService } from "../../../../../services/auth/auth.service";
import {FormBuilder,FormControl ,FormGroup, Validators} from '@angular/forms';
import { MatDialog } from '@angular/material/dialog';
import {Country} from '@angular-material-extensions/select-country'; 

@Component({
  selector: 'app-register-form',
  templateUrl: './register-form.component.html',
  styleUrls: ['./register-form.component.scss']
})
export class RegisterFormComponent implements OnInit {
  validationForm: FormGroup
  form: any = {}
  isSuccessful = false
  passwordFailed = false
  emailFailed = false
  isSignUpFailed = false
  errorMessage = '';
  check_email = '';
  check_pass = '';
  check_first_name = '';
  check_last_name = '';
  check_country = '';

  countryFormControl = new FormControl();
  countryFormGroup: FormGroup;

  roles: any[] = [
    {value: '2', name: 'Admin'},
    {value: '3', name: 'Chefe Trafego'},
    {value: '4', name: 'Camionista'}
  ];

  defaultValue: Country = {
    name: 'Portugal',
    alpha2Code: 'PT',
    alpha3Code: 'PRT',
    numericCode: '620',
    callingCode: '+351'
  };

  constructor(private authService: AuthService, public fb: FormBuilder, public dialog: MatDialog,private formBuilder: FormBuilder) {
    this.validationForm = fb.group({
      email: [null, [Validators.required, Validators.email]],
      role: [null, Validators.required],
      country:[null, Validators.required],
      password: [null, Validators.required],
      first_name:[null, Validators.required],
      last_name:[null, Validators.required],
      address:[null, Validators.required],
    });
  }
  
  get email() { return this.validationForm.get('email')}
  get password() { return this.validationForm.get('password')}
  get first_name() { return this.validationForm.get('first_name')}
  get last_name() { return this.validationForm.get('last_name')}
  get address() { return this.validationForm.get('address')}
  get role() { return this.validationForm.get('role').value }
  get country() { return this.validationForm.get('country').value.name}
  
  ngOnInit(): void {
   
  }

  onSubmit(): void {
    this.passwordFailed = false
    this.emailFailed = false
    this.isSignUpFailed = false
    this.authService.register(this.validationForm.value).subscribe(
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
        }else{
          this.emailFailed = true
        }
      }
    );
  }

  reloadPage(): void {
    window.location.reload();
  }


}
