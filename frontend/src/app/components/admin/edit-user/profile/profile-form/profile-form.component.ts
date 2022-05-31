import { Component, OnInit} from '@angular/core';
import { FormBuilder,FormControl ,FormGroup, Validators } from '@angular/forms';
import { MatDialog } from '@angular/material/dialog';
import { Country } from '@angular-material-extensions/select-country'; 
import { TokenService } from "../../../../../services/token/token.service";
import { UsersService } from "../../../../../services/user/users.service";
import { Router } from '@angular/router';

@Component({
  selector: 'app-profile-form',
  templateUrl: './profile-form.component.html',
  styleUrls: ['./profile-form.component.scss']
})
export class ProfileFormComponent implements OnInit {
  validationForm: FormGroup
  form: any = {}
  isSuccessful = false
  passwordFailed = false
  emailFailed = false
  isSignUpFailed = false
  errorMessage = '';
  check_email = '';
  check_first_name = '';
  check_last_name = '';
  check_address = '';
  check_country = '';
  name:any
  lastName:any
  addressbd:any
  emailbd:any

  countryFormControl = new FormControl();
  countryFormGroup: FormGroup;

  defaultValue: Country = {
    name: 'Portugal',
    alpha2Code: 'PT',
    alpha3Code: 'PRT',
    numericCode: '620',
    callingCode: '+351'
  };
 
  constructor(private readonly router: Router, private userService: UsersService, public fb: FormBuilder, public dialog: MatDialog, private tokenService: TokenService) {
    this.validationForm = fb.group({
      country:[null, Validators.required],
      first_name:[null, Validators.required],
      last_name:[null, Validators.required],
      address:[null, Validators.required],
    });
  }

  get address() { return this.validationForm.get('address')}
  get first_name() { return this.validationForm.get('first_name')}
  get last_name() { return this.validationForm.get('last_name')}
  get country() { return this.validationForm.get('country').value.name}
  
  ngOnInit(): void {
    const user = this.tokenService.getUser();
    this.name = user.Name
    this.lastName = user.LastName
    this.addressbd = user.Address
    this.emailbd = user.Email
  }

  onSubmit(): void {
    this.userService.updateUser(this.validationForm.value,this.emailbd).subscribe(
      data => {
        console.log(data)
        this.isSuccessful = true
        this.isSignUpFailed = false
        this.tokenService.signOut()
        this.router.navigate(['/login'])
      },
    );
  }

  reloadPage(): void {
    window.location.reload();
  }

}
