import {Component, OnInit} from '@angular/core';
import {UserService} from '../services/user/user.service';
import {Observable} from 'rxjs';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {

  isLoggedIn: Observable<boolean>;

  constructor(
    private userService: UserService
  ) {
  }

  ngOnInit(): void {
    this.isLoggedIn = this.userService.isLoggedIn;
  }

  onLogout(): void {
    this.userService.logout();
  }

}
