import { Component, OnInit } from '@angular/core'
import { Router, ActivatedRoute } from '@angular/router'
import { AuthService } from '../../services/auth/auth.service'
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-passrecover-form',
  templateUrl: './passrecover-form.component.html',
  styleUrls: ['./passrecover-form.component.scss']
})
export class PassrecoverFormComponent implements OnInit {
  passwordChange= '';
  recover = false
  token:any
  validationForm: FormGroup;
  constructor(private router:Router, private activeRoute:ActivatedRoute, private authService:AuthService , public fb: FormBuilder) { 
    this.validationForm = fb.group({
      confirme_password: [null, Validators.required],
      password: [null, Validators.required],
    });
  }

    get password() {
      return this.validationForm.get('password');
    }

    get confirmepassword() {
      return this.validationForm.get('confirme_password');
    }

  ngOnInit(): void {
    this.passwordChange= '';
    this.token = this.router.url.split("/passrecover/")[1]
    this.authService.passRecoverVerififyToken(this.token).subscribe(
      data => {
          this.recover = data.recover
      },
      err => {
        this.router.navigate(['/erro']).then(r =>
          this.reloadPage()
        );
      }
    );
  }
  
  onSubmit(): void {
    this.token = this.router.url.split("/passrecover/")[1]
    this.authService.passRecoverSetPassword(this.validationForm.value,this.token).subscribe(
      data => {
        this.passwordChange= 'true';
        this.router.navigate(['/login']).then(r =>
          this.reloadPage()
        );
      },
      err => {
        this.passwordChange= 'false';
      }
    );
  }
  reloadPage(): void {
    window.location.reload();
  }

}
