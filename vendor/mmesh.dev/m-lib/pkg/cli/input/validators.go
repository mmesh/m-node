package input

import (
	"errors"
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/crypto/ssh"
)

func ValidID(val interface{}) error {
	id := val.(string)

	if strings.ToLower(id) == "master" {
		return errors.New("reserved identifier")
	}

	r, err := regexp.MatchString(`^[a-z]{1}[a-z0-9-]+$`, id)
	if err != nil {
		return err
	}
	if !r {
		return errors.New("invalid identifier")
	}

	return nil
}

func ValidEmail(val interface{}) error {
	email := val.(string)

	r, err := regexp.MatchString(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`, email)
	if err != nil {
		return err
	}
	if !r {
		return errors.New("invalid email address")
	}

	return nil
}

func ValidNetwork(val interface{}) error {
	ipv4Addr, ipv4Net, err := net.ParseCIDR(val.(string))
	if err != nil {
		return err
	}

	if !ipv4Addr.Equal(ipv4Net.IP) {
		return fmt.Errorf("invalid network address %v", ipv4Addr)
	}

	cidrMask, _ := ipv4Net.Mask.Size()

	if cidrMask != 16 {
		return errors.New("only /16 networks are supported at the moment")
	}

	return nil
}

func ValidPassword(val interface{}) error {
	// var r bool
	// var err error

	pw := val.(string)

	if len(pw) < 10 {
		return errors.New("invalid password: length must be at least 10 characters")
	}

	/*
		r, err = regexp.MatchString(`[A-Z]`, pw)
		if err != nil {
			return err
		}
		if !r {
			return errors.New("invalid password")
		}

		r, err = regexp.MatchString(`[0-9]`, pw)
		if err != nil {
			return err
		}
		if !r {
			return errors.New("invalid password")
		}

		r, err = regexp.MatchString(`[a-z]`, pw)
		if err != nil {
			return err
		}
		if !r {
			return errors.New("invalid password")
		}
	*/

	return nil
}

func ValidIPNetCIDR(val interface{}) error {
	_, _, err := net.ParseCIDR(val.(string))
	if err != nil {
		return err
	}

	return nil
}

func ValidPort(val interface{}) error {
	port := val.(string)

	p, err := strconv.Atoi(port)
	if err != nil {
		return err
	}

	if p > 0 && p < 65535 {
		return nil
	}

	return errors.New("invalid port")
}

func ValidUint(val interface{}) error {
	n := val.(string)

	i, err := strconv.Atoi(n)
	if err != nil {
		return errors.New("invalid unsigned integer")
	}

	if i > 0 {
		return nil
	}

	return errors.New("invalid unsigned integer")
}

func ValidPrice(val interface{}) error {
	n := val.(string)

	p, err := strconv.Atoi(n)
	if err != nil {
		return errors.New("invalid price")
	}

	if p >= 0 {
		return nil
	}

	return errors.New("invalid price")
}

func ValidSSHPublicKey(val interface{}) error {
	pubKey := val.(string)

	if len(pubKey) == 0 {
		return errors.New("missing SSH public key")
	}

	if _, _, _, _, err := ssh.ParseAuthorizedKey([]byte(pubKey)); err != nil {
		return err
	}

	return nil
}
