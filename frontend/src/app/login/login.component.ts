import { Component, OnInit } from '@angular/core';
import { UserService} from '../services/user/user.service';
import {LoginInfo} from './loginInfo';
import {BehaviorSubject} from 'rxjs';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})

export class LoginComponent implements OnInit {

  username = '';
  password = '';
  isUsernameValid = true;
  error: any = null;

  pwdValue: string;
  usrValue: string;

  constructor(
    private loginService: UserService
  ) { }

  ngOnInit(): void {
    this.loginService.errorSubject.subscribe(errorMessage => {
      this.error = errorMessage;
    });
  }

  onKey(event: any, type: string): void {
    if (type === 'username') {
      this.username = event.target.value;
      this.validateUsername();
    } else if (type === 'password') {
      this.password = event.target.value;
    }
  }

  validateUsername(): void {
    const pattern = RegExp(/^[\w-.]*$/);
    if (pattern.test(this.username)) {
      this.isUsernameValid = true;
    } else {
      this.isUsernameValid = false;
    }
  }

  onSubmit(): void {
    if (this.isUsernameValid) {
      const user = { username: this.username, password: this.password };
      this.loginService.login(user);
    }
  }

}
